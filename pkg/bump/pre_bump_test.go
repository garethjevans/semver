package bump_test

import (
	"github.com/carvel-dev/semver/v4"
	"github.com/garethjevans/semver/pkg/bump"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PreBump", func() {
	var inputVersion semver.Version
	var b bump.PreBump
	var outputVersion semver.Version

	BeforeEach(func() {
		inputVersion = semver.Version{
			Major: 1,
			Minor: 2,
			Patch: 3,
		}

		b = bump.PreBump{}
	})

	JustBeforeEach(func() {
		outputVersion = b.Apply(inputVersion)
	})

	Context("when the prerelease is without version number", func() {
		BeforeEach(func() {
			b.Pre = "omega"
			b.PreWithoutVersion = true
		})

		Context("when the input is not a prerelease", func() {
			BeforeEach(func() {
				inputVersion.Pre = nil
			})

			It("bmps the prerelease without version number", func() {
				Expect(outputVersion).To(Equal(semver.Version{
					Major: 1,
					Minor: 2,
					Patch: 3,
					Pre: semver.PRVersion{
						{VersionStr: "omega"},
					},
				}))
			})
		})

		Context("when the input is a prerelease", func() {
			Context("when the bump is a different prerelease type", func() {
				BeforeEach(func() {
					inputVersion.Pre = semver.PRVersion{
						{VersionStr: "alpha"},
					}
				})

				It("bmps the prerelease without version number", func() {
					Expect(outputVersion).To(Equal(semver.Version{
						Major: 1,
						Minor: 2,
						Patch: 3,
						Pre: semver.PRVersion{
							{VersionStr: "omega"},
						},
					}))
				})
			})

			Context("when the bump is the same prerelease type", func() {
				BeforeEach(func() {
					inputVersion.Pre = semver.PRVersion{
						{VersionStr: "omega"},
						{VersionNum: 1, IsNum: true},
					}
				})

				It("bmps the prerelease without version number", func() {
					Expect(outputVersion).To(Equal(semver.Version{
						Major: 1,
						Minor: 2,
						Patch: 3,
						Pre: semver.PRVersion{
							{VersionStr: "omega"},
						},
					}))
				})
			})
		})
	})

	Context("when the version is a prerelease", func() {
		BeforeEach(func() {
			inputVersion.Pre = semver.PRVersion{
				{VersionStr: "alpha"},
				{VersionNum: 1, IsNum: true},
			}
		})

		Context("when the bump is the same prerelease type", func() {
			BeforeEach(func() {
				b.Pre = "alpha"
			})

			It("bumps the prerelease version number", func() {
				Expect(outputVersion).To(Equal(semver.Version{
					Major: 1,
					Minor: 2,
					Patch: 3,
					Pre: semver.PRVersion{
						{VersionStr: "alpha"},
						{VersionNum: 2, IsNum: true},
					},
				}))
			})

			It("does not mutate the input version", func() {
				Expect(inputVersion).To(Equal(semver.Version{
					Major: 1,
					Minor: 2,
					Patch: 3,
					Pre: semver.PRVersion{
						{VersionStr: "alpha"},
						{VersionNum: 1, IsNum: true},
					},
				}))
			})
		})

		Context("when the bump is a different prerelease type", func() {
			BeforeEach(func() {
				b.Pre = "beta"
			})

			It("bumps bumps to version 1 of the new prerelease type", func() {
				Expect(outputVersion).To(Equal(semver.Version{
					Major: 1,
					Minor: 2,
					Patch: 3,
					Pre: semver.PRVersion{
						{VersionStr: "beta"},
						{VersionNum: 1, IsNum: true},
					},
				}))
			})
		})
	})

	Context("when the version is not a prerelease", func() {
		BeforeEach(func() {
			inputVersion.Pre = nil
		})

		BeforeEach(func() {
			b.Pre = "beta"
		})

		It("bumps bumps to version 1 of the new prerelease type", func() {
			Expect(outputVersion).To(Equal(semver.Version{
				Major: 1,
				Minor: 2,
				Patch: 3,
				Pre: semver.PRVersion{
					{VersionStr: "beta"},
					{VersionNum: 1, IsNum: true},
				},
			}))
		})
	})
})
