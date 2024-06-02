Feature: Example scenario outline feature
  In order to test BDD with Godog
  As a developer
  I want to run scenarios with different sets of data

  Scenario Outline: Example scenario outline
    Given a precondition with <input>
    When an action is performed with <input>
    Then the outcome should be <output>

    Examples:
      | input | output  |
      | 1     | success |
      | 2     | success |
      | 3     | failure |
