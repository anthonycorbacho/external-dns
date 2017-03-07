// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package budgetsiface provides an interface to enable mocking the AWS Budgets service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package budgetsiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/budgets"
)

// BudgetsAPI provides an interface to enable mocking the
// budgets.Budgets service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Budgets.
//    func myFunc(svc budgetsiface.BudgetsAPI) bool {
//        // Make svc.CreateBudget request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := budgets.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockBudgetsClient struct {
//        budgetsiface.BudgetsAPI
//    }
//    func (m *mockBudgetsClient) CreateBudget(input *budgets.CreateBudgetInput) (*budgets.CreateBudgetOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockBudgetsClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type BudgetsAPI interface {
	CreateBudgetRequest(*budgets.CreateBudgetInput) (*request.Request, *budgets.CreateBudgetOutput)

	CreateBudget(*budgets.CreateBudgetInput) (*budgets.CreateBudgetOutput, error)

	CreateNotificationRequest(*budgets.CreateNotificationInput) (*request.Request, *budgets.CreateNotificationOutput)

	CreateNotification(*budgets.CreateNotificationInput) (*budgets.CreateNotificationOutput, error)

	CreateSubscriberRequest(*budgets.CreateSubscriberInput) (*request.Request, *budgets.CreateSubscriberOutput)

	CreateSubscriber(*budgets.CreateSubscriberInput) (*budgets.CreateSubscriberOutput, error)

	DeleteBudgetRequest(*budgets.DeleteBudgetInput) (*request.Request, *budgets.DeleteBudgetOutput)

	DeleteBudget(*budgets.DeleteBudgetInput) (*budgets.DeleteBudgetOutput, error)

	DeleteNotificationRequest(*budgets.DeleteNotificationInput) (*request.Request, *budgets.DeleteNotificationOutput)

	DeleteNotification(*budgets.DeleteNotificationInput) (*budgets.DeleteNotificationOutput, error)

	DeleteSubscriberRequest(*budgets.DeleteSubscriberInput) (*request.Request, *budgets.DeleteSubscriberOutput)

	DeleteSubscriber(*budgets.DeleteSubscriberInput) (*budgets.DeleteSubscriberOutput, error)

	DescribeBudgetRequest(*budgets.DescribeBudgetInput) (*request.Request, *budgets.DescribeBudgetOutput)

	DescribeBudget(*budgets.DescribeBudgetInput) (*budgets.DescribeBudgetOutput, error)

	DescribeBudgetsRequest(*budgets.DescribeBudgetsInput) (*request.Request, *budgets.DescribeBudgetsOutput)

	DescribeBudgets(*budgets.DescribeBudgetsInput) (*budgets.DescribeBudgetsOutput, error)

	DescribeNotificationsForBudgetRequest(*budgets.DescribeNotificationsForBudgetInput) (*request.Request, *budgets.DescribeNotificationsForBudgetOutput)

	DescribeNotificationsForBudget(*budgets.DescribeNotificationsForBudgetInput) (*budgets.DescribeNotificationsForBudgetOutput, error)

	DescribeSubscribersForNotificationRequest(*budgets.DescribeSubscribersForNotificationInput) (*request.Request, *budgets.DescribeSubscribersForNotificationOutput)

	DescribeSubscribersForNotification(*budgets.DescribeSubscribersForNotificationInput) (*budgets.DescribeSubscribersForNotificationOutput, error)

	UpdateBudgetRequest(*budgets.UpdateBudgetInput) (*request.Request, *budgets.UpdateBudgetOutput)

	UpdateBudget(*budgets.UpdateBudgetInput) (*budgets.UpdateBudgetOutput, error)

	UpdateNotificationRequest(*budgets.UpdateNotificationInput) (*request.Request, *budgets.UpdateNotificationOutput)

	UpdateNotification(*budgets.UpdateNotificationInput) (*budgets.UpdateNotificationOutput, error)

	UpdateSubscriberRequest(*budgets.UpdateSubscriberInput) (*request.Request, *budgets.UpdateSubscriberOutput)

	UpdateSubscriber(*budgets.UpdateSubscriberInput) (*budgets.UpdateSubscriberOutput, error)
}

var _ BudgetsAPI = (*budgets.Budgets)(nil)