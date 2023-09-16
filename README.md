# Newsletter Delivery System

## Project Overview

The Newsletter Delivery System is a comprehensive project comprising two distinct repositories, ["Publisher"](https://github.com/prayogatriady/newsletter-pub) and ["Subscriber"](https://github.com/prayogatriady/newsletter-sub), designed to streamline the process of managing and delivering newsletters to a broad audience. Leveraging cutting-edge technologies and best practices, this project demonstrates proficiency in software development, asynchronous communication, and effective email delivery.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Project Description](#Project-Description)
- [Authors](#authors)

## Getting Started

Explain how to install your Go project. Include any prerequisites, dependencies, or environment setup. Provide clear and concise instructions.

### Prerequisites

1. [Login](https://cloud.google.com/) to your GCP Account and [create](https://cloud.google.com/iam/docs/keys-create-delete) your `credential.json` file for Google Pub/Sub.
2. [Create](https://cloud.google.com/pubsub/docs/create-topic) Google Pub/Sub topic.
2. [Create](https://www.getmailbird.com/gmail-app-password/) your Gmail app password. (App password will be used in [the subscriber repository](https://github.com/prayogatriady/newsletter-sub))

### Installation

1. Clone the publisher repo

```sh
git clone https://github.com/prayogatriady/newsletter-pub.git
```

2. Move to the project root directory
```sh
cd newsletter-pub
```

3. Create `.env` file
```sh
PROJECT_ID=your-gcp-projectId
TOPIC_NAME=your-pubsub-topicName
```
4. Move your `credential.json` file to the project root directory

5. Get golang dependencies
```sh
go mod tidy
```

6. Run project
```sh
go run .
```

7. Send a HTTP POST request to this repository `localhost:8000/api/email/sendEmail`
```sh
{
    "recipient_list": [
        {
            "email": "example1@gmail.com",
            "name": "example1"
        },
        {
            "email": "example2@gmail.com",
            "name": "example2"
        },
        {
            "email": "example3@gmail.com",
            "name": "example3"
        },
        {
            "email": "example4@gmail.com",
            "name": "example4"
        }
    ]
}
```

#### The messages should be sent one by one asynchronously to your Google Pub/Sub topic.
#### Let's move to the [subscriber repository.](https://github.com/prayogatriady/newsletter-sub)

## Project Description

### Repository: Publisher

- **Description:** The "Publisher" repository serves as the backbone of the newsletter system. It fetches a curated list of email recipients from a RESTful API using the Gin web framework, ensuring efficient data retrieval and management.

- **Key Features:**
  - Utilizes Gin to create a RESTful API for email data retrieval.
  - Implements asynchronous publishing to Google Pub/Sub, optimizing performance and scalability.
  - Ensures data integrity and security during the publishing process.
  - Demonstrates proficiency in handling complex asynchronous workflows.

### Repository: Subscriber

- **Description:** The "Subscriber" repository is responsible for processing the published emails and delivering them to their respective recipients. It showcases the ability to handle email communication within a Go application.

- **Key Features:**
  - Efficiently receives and processes emails from Google Pub/Sub.
  - Utilizes the "gomail" package to send emails, allowing for customization and personalization of content.
  - Implements error handling and retries to ensure email delivery reliability.
  - Demonstrates strong attention to detail in email communication and management.

## Technologies and Skills Demonstrated

- **Golang (Go):** Proficient use of the Go programming language in both repositories, showcasing expertise in writing efficient, concurrent, and scalable code.

- **Microservices Architecture:** The project follows a microservices architecture, emphasizing the modular design of components for enhanced maintainability and scalability.

- **Asynchronous Communication:** Effective use of asynchronous communication between the Publisher and Subscriber repositories for improved performance and responsiveness.

- **RESTful API:** Development of a RESTful API for data retrieval, showcasing the ability to design and implement web services.

- **Google Pub/Sub:** Utilization of Google Pub/Sub for efficient and reliable message queuing and distribution.

- **Email Communication:** Implementation of email delivery using the "gomail" package, highlighting proficiency in email content customization and delivery.


## Project Impact

The Newsletter Delivery System demonstrates the capability to architect, develop, and manage complex, distributed systems. It underscores the commitment to delivering high-quality, reliable, and efficient software solutions. This project reflects the dedication to staying updated with the latest technologies and best practices in software development.

This project serves as a testament to the passion for creating innovative solutions and the dedication to excellence in software engineering. It is a tangible representation of the skills and experience in building robust, asynchronous, and scalable systems.

## Authors
- [Prayoga Triady](https://www.linkedin.com/in/prayogatriady/) - Backend Engineer