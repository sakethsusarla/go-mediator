# Mediator Example - Bazinga Smart Home System

## Prologue
The team at 2311 North Robles Avenue is working on building a Smart Home system. Howard Wolowitz, Sheldon Cooper, Leonard Hofstadter, and Rajesh Koothrappali are collaborating to create a centralized system that can control various smart devices like TVs, speakers, floor lamps, thermostats, and locks.

## Problem
Rajesh came up with an initial design where each device had to directly communicate with every other device and user. 

Sheldon, in his usual exasperated tone, pointed out several glaring issues with Rajesh's design:
1. _"This design is a textbook example of a tightly coupled system! Each device needs to know about every other device and user. It's like trying to solve a Rubik's cube blindfolded!"_
2. _"Adding new devices or users is a nightmare! You have to change things in multiple places. It's like trying to teach Penny physics!"_
3. _"As the number of devices and users increases, the complexity of the communication logic grows exponentially. It's approaching chaos theory levels of unmanageability!"_

## Solution
Sheldon proposed using the Mediator Design Pattern to solve these problems. In his typical didactic manner, he explained:

1. _"By introducing a mediator object, we can centralize the communication logic. It's like having a central command center that coordinates everything, rather than having each device act like a rogue agent."_
2. _"Devices and users will no longer need to know about each other. They only need to know about the mediator. It's like having a universal translator that everyone uses, instead of trying to learn every possible language."_
3. _"Adding new devices or users becomes trivial. You just register them with the mediator, and it handles the rest. It's like adding a new contact to your phone rather than rewriting the entire phonebook."_

Howard chimed in with his usual flair: "Alright, here's the deal. We create a `Mediator` interface that defines the communication methods. Then, we have a `SmartHomeMediator` that implements this interface and manages the interactions between `Device` objects and the `User` objects. Each `User` only talks to the `SmartHomeMediator`, making the system modular and easy to manage. It's like mission control for our devices and users!"
