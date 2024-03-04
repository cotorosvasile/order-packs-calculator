# Box Items Calculator

Web application that calculates the number of box items based on the provided pack sizes and quantity.

## Running the Application Locally

You can run this application on your local machine by following next steps:
1. Clone this repository on your local machine
2. Make sure you have docker installed on your machine
3. Run "docker-compose up -d" command in the root directory of this project
4. Open the `index.html` file in a web browser.
5. Enter the pack sizes in the "Pack Sizes" field. Pack sizes should be comma-separated numbers (e.g., "15, 54, 58, 68, 72, 91").
6. Enter the quantity in the "Quantity" field.
7. Click "Submit" to calculate the number of box items. The result will be displayed in a table below the form.

## Features

    If you run this calculation without submitting any values in the Pack Sizes field, you will notice that it will return some values. This is because the application contains cache implementation to store the default values for the boxes. Defaults are: 250, 500, 1000, 2000, 5000.


    There is a bdd folder under the /cmd directory which contains behavior tests. If you run "go test" command inside /cmd directory, it will bring up the application with docker-compose.yml and execute the Gherkin tests. 

