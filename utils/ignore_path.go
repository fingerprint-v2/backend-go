package utils

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/exp/slices"
)

func NewIgnorePathInstance(ignorePaths []string) (func(c *fiber.Ctx) bool, error) {
	regexList := make([]*regexp.Regexp, len(ignorePaths))
	for i, path := range ignorePaths {
		regex, err := regexp.Compile(path)
		if err != nil {
			return nil, err
		}
		regexList[i] = regex
	}

	return func(c *fiber.Ctx) bool {
		path := c.Path()
		return slices.ContainsFunc(regexList, func(regex *regexp.Regexp) bool {
			return regex.MatchString(path)
		})
	}, nil
}
