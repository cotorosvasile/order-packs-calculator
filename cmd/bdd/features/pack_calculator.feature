Feature: Calculate pack items

  Scenario Outline: 1. Successful calculation of pack items for default pack sizes
    Given system API is up and running
    When calculate POST request is send to "/calculator/calculate" with <itemsQty> in the request body
    Then the response code is 200
    And the response body contains <250>,<500>,<1000>,<2000>,<5000>

    Examples:
      | itemsQty        | 250  |500  |1000  |2000  |5000  |
      | 1               | 1    |0    |0     |0     |0     |
      | 250             | 1    |0    |0     |0     |0     |
      | 251             | 2    |0    |0     |0     |0     |
      | 501             | 1    |1    |0     |0     |0     |
      | 12001           | 1    |0    |0     |1     |2     |



  Scenario: 2. Missing items quantity in the request body
    Given system API is up and running
    When calculate POST request is send to "/calculator/calculate" with empty request body
    Then the response code is 400
    
