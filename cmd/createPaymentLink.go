/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"

	"github.com/manifoldco/promptui"
	"github.com/spf13/cobra"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/checkout/session"
)

var priceId string

// createPaymentLinkCmd represents the createPaymentLink command
var createPaymentLinkCmd = &cobra.Command{
	Use:   "createPaymentLink",
	Short: "Create a checkout session with Stripe.",
	Long: `Create a checkout session with Stripe.`,
	Run: func(cmd *cobra.Command, args []string) {

		stripeSecretKey, err := promptUser("Input your Stripe secret key", true)
		if err != nil {
			log.Fatalf("Unable to retrieve secret key %s", err)
		}
		stripe.Key = stripeSecretKey

		params := &stripe.CheckoutSessionParams{
		SuccessURL: stripe.String("https://example.com/success"),
		LineItems: []*stripe.CheckoutSessionLineItemParams{
			&stripe.CheckoutSessionLineItemParams{
			Price: stripe.String(priceId),
			Quantity: stripe.Int64(2),
			},
		},
		Mode: stripe.String(string(stripe.CheckoutSessionModePayment)),
		};
		result, err := session.New(params)
		if err != nil {
			log.Fatalf("Error while creating checkout session on Stripe: %s", err)
		}
		fmt.Println("Here is your payment link: " + result.URL)
	},
}

func promptUser(label string, hideEntered bool) (string, error) {
	validate := func(input string) error {
		if input == "" {
			return errors.New("input cannot be empty")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
		HideEntered: hideEntered,
	}

	result, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("Prompt failed %v\n", err)
	}

	return result, nil
}

func init() {
	rootCmd.AddCommand(createPaymentLinkCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createPaymentLinkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createPaymentLinkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	createPaymentLinkCmd.Flags().StringVarP(&priceId, "priceId", "p", "test", "Price id that has been set on the Stripe dashboard.")
}
