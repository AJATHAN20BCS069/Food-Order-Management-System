package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
	//make subTotalBill as global variable to make it easily accessible in case customer modifies her order.
	// subTotalBill means total bill but excluding taxes.
)

var subTotalBill float64

//make a map of customerOrder in which we will store the items ordered as "key" and no. of plates as "value".
var customerOrder = make(map[string]uint, 0)

func main() {
	var customerName string
	fmt.Println("What is your first name?")
	fmt.Scan(&customerName)

	fmt.Printf("%52s %s\n", "Vanakkam", customerName)
	fmt.Printf("%72s\n", "_/\\_ Aja's Restaurant _/\\_ ")
	fmt.Println()
	printMenu()
	var itemNumber uint
	var noOfPlates uint
	for {
		fmt.Println()
		fmt.Println("Enter '0' to exit.")
		fmt.Print("Enter the serial no. of the item to order: ")
		fmt.Scan(&itemNumber)
		if itemNumber == 0 {
			break
		}
		var choiceName string
		var itemPrice float64
		for index, item := range menu {
			if index+1 == int(itemNumber) {
				choiceName = item.itemName
				itemPrice = item.itemPrice
				break
			}
		}
		fmt.Printf("How many %v do you want: ", choiceName)
		fmt.Scan(&noOfPlates)
		if noOfPlates == 0 {
			continue
		} else {
			for index := range menu {
				if index+1 == int(itemNumber) {
					customerOrder[choiceName] += noOfPlates
					subTotalBill += itemPrice * float64(noOfPlates)
					break
				}
			}
			fmt.Printf("\nYou just ordered %v %v which amounts to ₹%v.\n", noOfPlates, choiceName, itemPrice*float64(noOfPlates))
			orderTillNow()
		}
		fmt.Println()
	}
	fmt.Println()
	billDisplayText := "************************************* Generating Bill *************************************"
	for _, element := range billDisplayText {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 15)
	} //just displays that "generating bill" in a fancy manner.
	fmt.Println()
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	fmt.Printf(" %-30s %-20s %-20s %-20s\n", "Item Name", "Price(₹)", "Quantity", "Total Price(₹)")
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	printOrderData()
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	fmt.Printf("%47s", " ")
	for _, element := range "Sub Total(excluding GST): ₹" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Printf("%.2f\n", subTotalBill)
	func() {
		for {
			var isOrderOK string
			fmt.Println("Do you want to change your order?[y/n]")
			fmt.Scan(&isOrderOK)
			if isOrderOK != "y" {
				return
			}
			var serialNo uint
			var modifyType uint
			fmt.Println("Please enter the respective no. to proceed further: ")
			fmt.Println("Press '1' to update item quantity.")
			fmt.Println("Press '2' to delete an item from the order list.")
			fmt.Println("Press '3' to add item(s) in the order list.")
			fmt.Scan(&modifyType)
			switch modifyType {
			case 1:
				printMenu()
				fmt.Println("Please enter the S.No. of the item to be updated: ")
				fmt.Scan(&serialNo)
				updateQuantity(serialNo)
			case 2:
				printMenu()
				fmt.Println("Please enter the S.No. of the item to be deleted: ")
				fmt.Scan(&serialNo)
				delFromOrder(serialNo)
			case 3:
				insertIntoOrder()
			default:
				return
			}
			displayGeneratingBill()
			generateBill()
		}
	}()
	for _, element := range "Here is your final bill:-" {
		fmt.Printf("%c", element)
		time.Sleep(time.Millisecond * 50)
	}
	fmt.Println()
	fmt.Printf("\n%52s\n", "Aja's Restaurant")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s\n", strings.Repeat("*", 91))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%86s\n", "Balaji Nagar,Eathakoil Road,Aundipatty-Theni District,625512,Tamil Nadu,India")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%50s\n", "Tel: 6379857346")
	fmt.Printf("%60s\n\n", "Email: ajasrestaurantorder@gmail.com")
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%s", strings.Repeat("-", 42))
	fmt.Printf("%s", "INVOICE")
	fmt.Printf("%s\n", strings.Repeat("-", 42))
	time.Sleep(time.Millisecond * 200)
	rand.Seed(time.Now().Unix())
	fmt.Printf(" Ticket No: %d\n", rand.Intn(550)+1)
	fmt.Printf(" Date: %v\n", time.Now().Local().Format("06-Jan-02"))
	fmt.Printf(" Time: %v", time.Now().Local().Format("3:4 PM"))
	time.Sleep(time.Millisecond * 200)
	generateBill()
	tax := 18 * subTotalBill / (100)
	grandTotal := subTotalBill + tax
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: ₹%.2f\n", "GST", tax)
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	time.Sleep(time.Millisecond * 200)
	fmt.Printf("%71s: ₹%.2f\n", "Grand Total", grandTotal)
	fmt.Printf("+%s+\n", strings.Repeat("-", 90))
	{
		var customerName string = customerName
		fmt.Println()
		fmt.Printf("%17s", " ")
		fmt.Printf("_/\\_ Thank you %v for visiting Aja's Restaurant _/\\_\n", customerName)
		fmt.Printf("%55s\n", "We hope to see you again!")
		fmt.Printf("%51s\n", "Have a nice day! ")
	}
}
