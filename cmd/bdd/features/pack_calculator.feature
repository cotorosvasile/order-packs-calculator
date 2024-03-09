Feature: Calculate pack items

  Scenario: 1. Missing items quantity in the request body
    Given system API is up and running
    When calculate POST request is send to "/calculator/calculate" with empty request body
    Then the response code is 400

  Scenario Outline: 2. Success flow. Result contains least items than least boxes
    Given system API is up and running
    When calculate POST request is send to "/calculator/calculate" with <boxSizes> and <itemsQty> in the request body
    Then the response code is 200
    And the response body contains <packConfig>

    Examples:
    | boxSizes                    | itemsQty  | packConfig                                     |
    | "[250,500,1000,2000,5000]"  | 1         | '{"box_items": {"250": 1}}'                    |
    | "[250,500,1000,2000,5000]"  | 250       | '{"box_items": {"250": 1}}'                    |
    | "[250,500,1000,2000,5000]"  | 251       | '{"box_items": {"500": 1}}'                    |
    | "[250,500,1000,2000,5000]"  | 501       | '{"box_items": {"250": 1,"500": 1}}'           |
    | "[250,500,1000,2000,5000]"  | 12001     | '{"box_items": {"2000": 1,"250": 1,"5000": 2}}'|
    | "[5,12]"                    | 14        | '{"box_items": {"5": 3}}'                      |
