package cmd_test

import (
	"bytes"
	command "github.com/ddubson/omni/cmd"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"io/ioutil"
)

var _ = Describe("whoami", func() {
	It("should return signed in user", func() {
		whoAmICmd := command.NewWhoAmICmd()
		// Explicitly set 0 args for a command that has no args, otherwise test args are provided to the command which
		// throws an error
		whoAmICmd.SetArgs([]string{})
		buffer := bytes.NewBufferString("")

		whoAmICmd.SetOut(buffer)
		err := whoAmICmd.Execute()
		if err != nil {
			Fail(err.Error())
		}
		output, _ := ioutil.ReadAll(buffer)
		Expect(string(output)).To(ContainSubstring("You are signed in as "))
	})
})
