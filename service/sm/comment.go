package sm

import (
	"context"
	model "github.com/ctreminiom/go-atlassian/pkg/infra/models"
)

type CommentConnector interface {

	// Gets returns all comments on a customer request.
	//
	// No permissions error is provided if, for example, the user doesn't have access to the service desk or request,
	//
	// the method simply returns an empty response.
	//
	// GET /rest/servicedeskapi/request/{issueIdOrKey}/comment
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#get-request-comments
	Gets(ctx context.Context, issueKeyOrID string, public bool, expand []string, start, limit int) (*model.RequestCommentPageScheme, *model.ResponseScheme, error)

	// Get returns details of a customer request's comment.
	//
	// GET /rest/servicedeskapi/request/{issueIdOrKey}/comment/{commentId}
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#get-request-comment-by-id
	Get(ctx context.Context, issueKeyOrID string, commentID int, expand []string) (*model.RequestCommentScheme, *model.ResponseScheme, error)

	// Create creates a public or private (internal) comment on a customer request, with the comment visibility set by public.
	//
	// POST /rest/servicedeskapi/request/{issueIdOrKey}/comment
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#create-request-comment
	Create(ctx context.Context, issueKeyOrID, body string, public bool) (*model.RequestCommentScheme, *model.ResponseScheme, error)

	// Attachments  returns the attachments referenced in a comment.
	//
	// GET /rest/servicedeskapi/request/{issueIdOrKey}/comment/{commentId}/attachment
	//
	// https://docs.go-atlassian.io/jira-service-management-cloud/request/comments#get-comment-attachments
	Attachments(ctx context.Context, issueKeyOrID string, commentID, start, limit int) (*model.RequestAttachmentPageScheme, *model.ResponseScheme, error)
}
