package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func main() {
	var productID int
	var action string

	// Ask for product ID
	fmt.Print("Enter product ID: ")
	_, err := fmt.Scan(&productID)
	if err != nil {
		fmt.Println("Error reading product ID:", err)
		return
	}

	// Ask for action type
	fmt.Print("Enter action (prediction/comparision): ")
	_, err = fmt.Scan(&action)
	if err != nil {
		fmt.Println("Error reading action:", err)
		return
	}

	if action == "comparision" {
		predictions, err := Comparision_Prediction(productID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Predictions:", predictions)
	} else if action == "prediction" {
		predictions, err := RunPrediction(productID)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("Predictions:", predictions)
	} else {
		fmt.Println("Invalid action. Please enter 'prediction' or 'comparision'.")
	}
}
func Comparision_Prediction(productID int) ([]int, error) {
	// Construct the command to call Python script with arguments
	cmd := exec.Command("python3", "Comparision_Prediction.py", fmt.Sprintf("%d", productID))

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error executing script: %v", err)
	}

	// Parse JSON output
	var predictions []int
	if err := json.Unmarshal(output, &predictions); err != nil {
		return nil, fmt.Errorf("error parsing JSON output: %v", err)
	}

	return predictions, nil
}

func RunPrediction(productID int) ([]int, error) {
	// Construct the command to call Python script with arguments
	cmd := exec.Command("python3", "Prediction.py", fmt.Sprintf("%d", productID))

	// Run the command and capture the output
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error executing script: %v", err)
	}

	// Parse JSON output
	var predictions []int
	if err := json.Unmarshal(output, &predictions); err != nil {
		return nil, fmt.Errorf("error parsing JSON output: %v", err)
	}

	return predictions, nil
}
