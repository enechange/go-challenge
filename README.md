ENECHANGE Go Developer Challenge
====

### Why do we ask you to do this challenge?

As part of the ENECHANGE interview process, we are asking you to work on a small coding exercise to help us understand your skills, and give you an idea of the work you would be doing with us.

### Objectives

Create an endpoint that meets the following requirements using the template.

Please document any technical decisions, trade-offs, problems, etc., in REPORT.md (you may write in either Japanese or English).

#### Requirements:
Please create an endpoint for searching chargers within a specified range, based on the interface specifications provided in the PDF document.
Use the data stored in the provided CSV files by importing it into the database.

-  Specification:
   - [Go-Challenge Interface Specification Document](./Go-Challenge%20Interface%20Specification%20Document.pdf)

-  Data:
    - [locations.csv](./resources/locations.csv)
    - [evses.csv](./resources/evses.csv)

#### Language / Libraries:
- Language: GO
- Libraries:
  - Web Application Framework: [Gin](https://gin-gonic.com/)
  - ORM: [GORM](https://gorm.io/)

#### Template:
We have prepared a template that sets up a server and connects to a database.
Please add code in the pkg directory to implement the endpoint.

#### Steps
1. Please fork this repository.
2. Create a new branch. (The branch name should be `challenge/YOURACCOUNT` where `YOURACCOUNT` is your Github Account name, for example, `challenge/shirakia`)
3. Proceed with your implementation using provided template.
4. After completion, create a Pull Request and please provide us with the URL.
5. You may continue to commit to your branch even after you have provided us with the URL of the Pull Request.