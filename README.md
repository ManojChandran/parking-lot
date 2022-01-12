# Design a General Parking-lot
Design a parking lot management software, which help to manage a parking space.

# Design life cycle
## Feature list
1) Car parking space.                                                            </br>
2) Car parking price.                                                            </br>
3) Van parking space.                                                            </br>
4) Van parking price.                                                            </br>
5) Motor bikes parking space.                                                    </br>
6) Motor bike parking price.                                                     </br>
7) Know parking space availability.                                              </br>
8) Get ticket for parking vehicle.                                               </br>
9) Tag ticket to corresponding parking type.                                      </br>
10) Ticket should have entry time and vehicle type.                              </br>
11) Payment calculated on exit, based on entry time and vehicle parking price.   </br>

## Use case Diagram
## Break up problem

## Requirement
Designed parking should have 

* 10k Parking capacity
* 4 Entrance 
* Collect ticket at entances
* 3 Types of parking (car, van and motor cycle)
* System should not allow more vehicles than Parking capacity
* Hourly based parking calculation

## Domain Analysis
## Preliminary Design
## Implementation 
## Delivery

# Entity
1. Parking
    - Parking Company name
    - Parking Capacity Car
    - Parking Capacity Van
    - Parking Capacity Motor cycle
2. Parking Slot
    - Parking Type
    - Parking Slot number
    - Parking Status
3. Parking Ticket
    - Vehicle number
    - Vehicle type
    - Vehicle Entry time
    - Vehicle Exit time
4. Parking Gate
    - Entry ticket
    - Exit Bill

# Use Case 1
### ID
PARK0001
### Case Title
Vehicle Type
### Description
Parking lot owner should decide on types of vehicle, which allowed to park in his or her facility.
### Primary actor
owner
### Pre Condition
None
### Post Condition

### Main Success scenario
Should have list of allowed vehicle type
### Extensions
### Frequency of use
### Status
### Owner
### Priority

# Use Case 2
### ID
PARK0002
### Case Title
Parking ticket
### Description
User should get a parking ticket at entrance. Parking ticket should have a unique id, entry time Vehicle type and Vehicle number printed. Entery time should be a of ISO 8601 format. Timestamp denotes the entry time date of vehicle and it should be current date and time.
### Primary actor
Car owner
### Pre Condition
Allowed vehicle types should be defined.
### Post Condition

### Main Success scenario
### Extensions
### Frequency of use
### Status
### Owner
### Priority