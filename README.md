# Agnostic AI (WIP - Golang Noob Here)

This rest application is built in **Go** using the **Fiber** web framework to provide a flexible and agnostic API for interacting with various AI models. The goal is to allow users to send requests, which can be processed by any AI model, whether it's a machine learning model, natural language processing model, or any other type of AI.

## Features

- **Flexible AI Integration**: The application is designed to work with any AI model, either by invoking APIs or interacting directly with pre-trained models.
- **RESTful API**: Offers a clean and simple HTTP interface to communicate with AI systems.
- **AI Agnosticism**: The API does not make assumptions about the type of AI model being used. It is up to the user to specify the model they wish to interact with.
- **Support for multiple AI models**: The system can easily be extended to support additional models in the future.

## Supported AI Models

- **OpenAI GPT**
  - [x] Text Generation
  - [x] Image Generation
  - [x] Audio Transcription
- **Gemini**
  - [x] Text Generation
  - [ ] Image Generation
  - [ ] Audio Transcription
- **Claude**
  - [x] Text Generation
- **Ollama**
  - [x] Text Generation
- **Groq**
  - [x] Text Generation
  - [x] Audio Transcription
- **Huggingface**
  - [ ] Text Generation
  - [ ] Image Generation
  - [ ] Audio Transcription

## Requirements

- **Go** (1.18+ recommended)
- **Fiber** web framework
- An external AI service or a locally integrated AI model.

## Installation

### Prerequisites

Ensure you have Go installed on your system. If not, you can install Go from [here](https://golang.org/dl/).

### Clone the Repository

git clone <https://github.com/endrureza/agnostic-ai.git>
cd agnostic-ai

### Install Dependencies

go mod tidy

### Run the Application

To start the application, use the following command:

go run main.go

The application will start a web server on port `5000`.

---

## API Endpoints

### 1. **POST /chat**

This endpoint to generate chat response.

#### Request Body

{
  "provider": "openai",
  "model": "gpt-4o-mini",
  "messages": [
    {
        "role": "user",
        "content": "Hello, how are you?"
    }
  ]
}

#### Response

{
  "text": "I'm fine thank you"
}

#### Example Usage

curl -X POST <http://localhost:5000/chat> \
  -H "Content-Type: application/json" \
  -d '{
    "provider": "openai",
    "model": "gpt-4o-mini",
    "messages": [
      {
        "role": "user",
        "content": "Hello, how are you?"
      }
    ]
  }'
