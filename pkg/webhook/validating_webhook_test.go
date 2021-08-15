package webhook_test

import (
	"context"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/snorwin/k8s-generic-webhook/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"
)

var _ = Describe("Validating Webhook", func() {
	Context("ValidateFuncs", func() {
		It("should by default allow all", func() {
			result := (&webhook.ValidateFuncs{}).ValidateCreate(context.TODO(), admission.Request{})
			Ω(result.Allowed).Should(BeTrue())
			result = (&webhook.ValidateFuncs{}).ValidateUpdate(context.TODO(), admission.Request{})
			Ω(result.Allowed).Should(BeTrue())
			result = (&webhook.ValidateFuncs{}).ValidateDelete(context.TODO(), admission.Request{})
			Ω(result.Allowed).Should(BeTrue())
		})
		It("should use defined functions", func() {
			result := (&webhook.ValidateFuncs{
				CreateFunc: func(ctx context.Context, _ admission.Request) admission.Response {
					return admission.Denied("")
				},
			}).ValidateCreate(context.TODO(), admission.Request{})
			Ω(result.Allowed).Should(BeFalse())
		})
	})
})
