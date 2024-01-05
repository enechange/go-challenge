ENECHANGE Go Developer Challenge
====

### Why do we ask you to do this challenge?

As part of the ENECHANGE interview process, we are asking you to work on a small coding exercise to help us understand your skills, and give you an idea of the work you would be doing with us.

### Objectives

Create an endpoint that meets the following requirements using the specified language, library, and software design model.

Please document any technical decisions, trade-offs, problems, etc., in REPORT.md (you may write in either Japanese or English).

#### Requirements:
-  Specification:
   - [OCPI 2.2](https://evroaming.org/app/uploads/2020/06/OCPI-2.2-d2.pdf) / Locations module / Sender Interface / GET Method (P.48)
      - There is no need to read the entire OCPI document.
      - Implementation for authentication (e.g., OCPI Credentials module) is not required.

- You are free to decide the endpoint path.

- The response should be as follows:
  - Return 3 Locations if there is no specified filtering.
  - Each Location should have between 1 and 3 EVSEs.
  - Each EVSE should have one Connector.
  - Include only the following specified items in the response and process any unnecessary items according to OCPI Cardinality (P.9):
    - The values of each item can be anything as long as they match the specification.
    - Location:
      - country_code
      - party_id
      - id
      - publish
      - name
      - address
      - city
      - country
      - coordinates
      - evses
      - time_zone
      - opening_times
          - twentyfourseven
          - regular_hours
      - last_updated
    - EVSE:
      - uid
      - evse_id
      - status
      - connectors
      - last_updated
    - Connector:
      - id
      - standard
      - format
      - power_type
      - max_voltage
      - max_amperage
      - last_updated

#### Language / Library / Software Design Model:
- Language: GO
- Libraries:
  - Web Application Framework: [Gin](https://gin-gonic.com/)
  - ORM: [GORM](https://gorm.io/)
- Software Design Model: MVC

#### Steps
1. Please fork this repository.
2. Create a new branch. (The branch name should be challenge/YOURNAME where YOURNAME can be your real name or your Github Account name, for example, challenge/shirakia)
3. Create a development directory directly under the repository. (e.g., shirakia)
4. Proceed with your implementation inside this directory.
5. After completion, create a Pull Request and please provide us with the URL.
6. You may continue to commit to your branch even after you have provided us with the URL of the Pull Request.