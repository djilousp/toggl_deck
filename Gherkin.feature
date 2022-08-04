
Feature: Create Deck:

    Scenario: Create Full Shuffled Deck
        Given a player wants to generate a deck 
        When he chooses a shuffled deck
        And no cards are wanted
        Then a new 52-shuffled deck is returned
    
    Scenario: Create Partial Unshuffled Deck
        Given a player wants to generate a deck 
        When he chooses to get an unshuffled deck
        And  a list of cards is specified "AS,KD,AC,2C,KH"
        Then a new partiel-unshuffled deck is created from the list of cards 
        

Feature: Open Deck: 

    Background:
        Given a player who wants to play a card game 

    Scenario: Deck should befor a given :
        When the ID "a251071b-662f-44b6-ba11-e24863039c59" is entered 
        Then the deck with ID "a251071b-662f-44b6-ba11-e24863039c59" is returned

    Scenario: Error When Invalid or not passed ID:
        When an ID is not passed or Invalid 
        Then the deck is not found and an error is returned

Feature: Draw Card: 

    Background:
        Given a player who wants to draw cards 

    Scenario: Cards with X amount should be drawn from Deck :
        When the ID "a251071b-662f-44b6-ba11-e24863039c59" is entered 
        And a count of 6 to be drawn
        Then 6 cards are returned

    Scenario: Error When Invalid or not passed ID:
        When an ID is not passed or Invalid 
        And a count of 6 to be drawn
        Then the deck is not found and an error is returned