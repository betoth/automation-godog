Feature: Fetch Data from Public API

  Scenario: Retrieve data from a public API
    Given I make a GET request to the public API
    Then the response status code should be 200
