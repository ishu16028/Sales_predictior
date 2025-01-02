# Sales Predictor

This project predicts sales for the next 12 weeks using time series data as input and the SARIMAX model for forecasting.

## Project Description

The Sales Predictor is designed to forecast sales for a given period using historical time series data. This can help in planning and decision-making for inventory management, marketing strategies, and financial forecasting.

## Installation and Usage Instructions

1. **Unzip the dataset:**
   - Unzip the `1.zip` folder as it contains the dataset.
   - Place a copy of the dataset in the parent directory of the project.

2. **Dependencies:**
   - Ensure you have Python and Jupyter Notebook installed.
   - Install the required Python packages:
     ```sh
     pip install -r requirements.txt
     ```

3. **Running the Project:**
   - Open the Jupyter Notebook.
   - Load the dataset and run the cells to train the SARIMAX model and predict sales for the next 12 weeks.

## Methodology

The project uses the SARIMAX (Seasonal AutoRegressive Integrated Moving Average with eXogenous regressors) model to predict future sales based on historical data. This model is well-suited for time series forecasting with seasonal patterns.

## Code Summary

### Data Loading and Transformation:

- Load data into a DataFrame.
- Parse dates and extract relevant time features (year and week_of_year).

### Aggregation:

- Aggregate sales data into weekly quantities.
- Include exogenous variables (discount_applied, product_stock) for enhanced modeling.

### SARIMAX Model Training:

- Train a SARIMAX model for a specified product ID.
- Handle cases where no valid data exists for the product.

### Forecasting and Visualization:

- Forecast future sales using the trained model.
- Visualize historical and forecasted sales.
