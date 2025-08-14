# RabbitMQ Go Publisher/Subscriber Example

This project demonstrates a basic messaging system using RabbitMQ with a Go publisher that sends messages and a Go subscriber that receives them.

## 🚀 How to Run the Project

### Prerequisites
- Docker installed
- Docker Compose installed

### Setup Instructions

1. **Clone the repository**:
   ```bash
   git clone https://github.com/VladimirAzanza/rabbitmq_small_easy_example.git
   cd rabbitmq_small_easy_example/src
   ```

2. **Build and run the containers**:
    ```bash
    docker-compose up --build
    ```

### Accessing Services

1. RabbitMQ Management UI:
    Open http://localhost:15672 in your browser
    Credentials: guest / guest

Publisher Service:
Sends messages to the queue every 1 second
Logs output to the console

Subscriber Service:
Receives messages from the queue
Logs incoming messages to the console