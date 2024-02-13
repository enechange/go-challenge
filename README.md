ENECHANGE Go Developer Challenge
====

### Why do we ask you to do this challenge?

As part of the ENECHANGE interview process, we are asking you to work on a small coding exercise to help us understand your skills, and give you an idea of the work you would be doing with us.

### Objectives

Create an endpoint that meets the following requirements using the template.

Please document any technical decisions, trade-offs, problems, etc., in REPORT.md (you may write in either Japanese or English).

#### Requirements:
Please create an endpoint to search for chargers within a specified range, according to the interface specifications outlined in the provided PDF document.

You can use the provided CSV files as sample data by importing them into the database.

While writing tests is not mandatory, please be mindful to design for testability.

- Specification:
   - [Go-Challenge Interface Specification Document](./Go-Challenge%20Interface%20Specification%20Document.pdf)

- CSV Files:
    - [locations.csv](./sample/locations.csv)
    - [evses.csv](./sample/evses.csv)

#### Language / Libraries:
- Language: GO
- Libraries:
  - Web Application Framework: [Gin](https://gin-gonic.com/)
  - ORM: [GORM](https://gorm.io/)

#### Template:
A template has been prepared in the codebase that sets up a server and prepares the connection to a database.

Please add code to implement the endpoint.

#### Steps
1. Please fork this repository.
2. Create a new branch. (The branch name should be `challenge/YOURACCOUNT` where `YOURACCOUNT` is your Github Account name, for example, `challenge/shirakia`)
3. Proceed with your implementation using provided template.
4. After completion, create a Pull Request and please provide us with the URL.
5. You may continue to commit to your branch even after you have provided us with the URL of the Pull Request.