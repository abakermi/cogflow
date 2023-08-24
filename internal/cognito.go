package internal

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

// CognitoClient is a struct that holds the Cognito service client.
type CognitoClient struct {
	svc *cognitoidentityprovider.CognitoIdentityProvider
}

// User holds user attributes, username, and status.
type User struct {
	Username       string
	UserAttributes []*cognitoidentityprovider.AttributeType
	Status         string
}


// NewCognitoClient creates a new CognitoClient instance.
func NewCognitoClient() (*CognitoClient, error) {
	awsProfile := GetAWSProfile()
	region := GetRegion()

	// If the profile is not provided, set it to default
	if awsProfile == "" {
		awsProfile = "default"
	}
	var sess *session.Session
	var err error

	// If the region is provided, create a session with the specified region
	if region != "" {
		sess, err = session.NewSession(&aws.Config{
			Region: aws.String(region),
		})
	} else {
		// Otherwise, use the session with the profile
		sess, err = session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
			Profile:           awsProfile,
		})
	}

	if err != nil {
		LogError("Error creating session:", err)
		return nil, err
	}

	svc := cognitoidentityprovider.New(sess)

	return &CognitoClient{
		svc: svc,
	}, nil
}

// GetUserByID retrieves user details by their Cognito user ID.
func (c *CognitoClient) GetUserByID(userID string) (*User, error) {
	userPoolID := GetPoolID()
	input := &cognitoidentityprovider.AdminGetUserInput{
		UserPoolId: aws.String(userPoolID),
		Username:   aws.String(userID),
	}

	result, err := c.svc.AdminGetUser(input)
	if err != nil {
		println("getting user:", err)
		return nil, err
	}
	user := &User{
		Username:       *result.Username,
		UserAttributes: result.UserAttributes,
		Status:         *result.UserStatus,
	}
	return user, nil
}
// EnableUser enables a user in Amazon Cognito by their Cognito user ID.
func (c *CognitoClient) EnableUser(userID string) error {
	userPoolID := GetPoolID()
	input := &cognitoidentityprovider.AdminEnableUserInput{
		UserPoolId: aws.String(userPoolID),
		Username:   aws.String(userID),
	}

	_, err := c.svc.AdminEnableUser(input)
	if err != nil {
		println("enabling user:", err)
		return err
	}

	return nil
}

// DisableUser disables a user in Amazon Cognito by their Cognito user ID.
func (c *CognitoClient) DisableUser(userID string) error {
	userPoolID := GetPoolID()
	input := &cognitoidentityprovider.AdminDisableUserInput{
		UserPoolId: aws.String(userPoolID),
		Username:   aws.String(userID),
	}

	_, err := c.svc.AdminDisableUser(input)
	if err != nil {
		println("disabling user:", err)
		return err
	}

	return nil
}

// DeleteUser deletes a user from Amazon Cognito by their Cognito user ID.
func (c *CognitoClient) DeleteUser(userID string) error {
	userPoolID := GetPoolID()
	input := &cognitoidentityprovider.AdminDeleteUserInput{
		UserPoolId: aws.String(userPoolID),
		Username:   aws.String(userID),
	}

	_, err := c.svc.AdminDeleteUser(input)
	if err != nil {
		println("deleting user:", err)
		return err
	}
	return nil
}


// ListGroups retrieves a list of all available groups.
func (c *CognitoClient) ListGroups() ([]string, error) {
	userPoolID := GetPoolID()

	input := &cognitoidentityprovider.ListGroupsInput{
		UserPoolId: aws.String(userPoolID),
	}

	result, err := c.svc.ListGroups(input)
	if err != nil {
		return nil, err
	}

	groups := make([]string, len(result.Groups))
	for i, group := range result.Groups {
		groups[i] = *group.GroupName
	}

	return groups, nil
}

// CreateGroup creates a new group in the Cognito user pool.
func (c *CognitoClient) CreateGroup(groupName string) error {
	userPoolID := GetPoolID()

	input := &cognitoidentityprovider.CreateGroupInput{
		GroupName:   aws.String(groupName),
		UserPoolId:  aws.String(userPoolID),
		Description: aws.String("Description for " + groupName),
	}

	_, err := c.svc.CreateGroup(input)
	if err != nil {
		return err
	}

	return nil
}

// DeleteGroup deletes a group by its name.
func (c *CognitoClient) DeleteGroup(groupName string) error {
	userPoolID := GetPoolID()

	input := &cognitoidentityprovider.DeleteGroupInput{
		GroupName:  aws.String(groupName),
		UserPoolId: aws.String(userPoolID),
	}

	_, err := c.svc.DeleteGroup(input)
	if err != nil {
		return err
	}

	return nil
}
