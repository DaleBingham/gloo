package nsselect

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

var _ = Describe("Generate Options", func() {
	nsrMap := make(NsResourceMap)
	nsrMap["ns1"] = &NsResource{
		Meshes:  []string{"m1", "m2"},
		Secrets: []string{"s1", "s2"},
		MeshesByInstallNs: []core.ResourceRef{
			core.ResourceRef{
				Name:      "m1",
				Namespace: "ns1",
			},
		},
	}
	nsrMap["ns2"] = &NsResource{
		Meshes:  []string{},
		Secrets: []string{"s3"},
		MeshesByInstallNs: []ResourceRef{
			ResourceRef{
				Name:      "m2",
				Namespace: "ns1",
			},
		},
	}

	It("should create the correct Secret options and map", func() {
		genOpts, resMap := generateCommonResourceSelectOptions("secret", nsrMap)
		Expect(genOpts).To(Equal([]string{
			"ns1, s1",
			"ns1, s2",
			"ns2, s3",
		}))
		expectedMap := make(ResMap)
		expectedMap["ns1, s1"] = ResSelect{
			displayName:      "s1",
			displayNamespace: "ns1",
			resourceRef: core.ResourceRef{
				Name:      "s1",
				Namespace: "ns1",
			},
		}
		expectedMap["ns1, s2"] = ResSelect{
			displayName:      "s2",
			displayNamespace: "ns1",
			resourceRef: core.ResourceRef{
				Name:      "s2",
				Namespace: "ns1",
			},
		}
		expectedMap["ns2, s3"] = ResSelect{
			displayName:      "s3",
			displayNamespace: "ns2",
			resourceRef: core.ResourceRef{
				Name:      "s3",
				Namespace: "ns2",
			},
		}
		Expect(resMap).To(Equal(expectedMap))
	})
})
