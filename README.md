# Ping - Stock Ticker Service 
This repository contains a Go web service that retrieves and displays stock data for a specific symbol over a set number of days, along with a Kubernetes deployment configuration to deploy the service on Minikube or a Kubernetes cluster. 

### Contents
- `stock_ticker/`: Contains the Go source code and HTML templates for the web service.
- `kubernetes/`: Contains Kubernetes manifests for deploying the service, setting up a horizontal pod autoscaler (HPA), a NodePort service, a ConfigMap for environment variables, and a Secret for the API key.
- `Dockerfile`: In the root directory `Ping/` for building the Docker image if needed.

### Prerequisites
Ensure you have the following installed:
- **Go** (Golang)
- **Docker**
- **Minikube** (or Kubernetes)
- **kubectl** CLI

### Setup Instructions
These instructions are for Minikube on a Windows host. If you are using a different Kubernetes setup, you may need to adjust the `service.yaml` for your environment. The current setup uses a NodePort.

#### Step 1: Clone the Repository
Clone this repository to your local machine:
```bash
git clone https://github.com/giselle737/Ping.git
cd Ping
```

#### Step 2: Start Minikube and Docker
Start Minikube with the Docker driver:
```powershell
minikube start --driver=docker
```

Verify that Docker is running and connected to Minikube:
```powershell
minikube status
```

#### Step 3: Deploy the Kubernetes Manifests
Navigate to the `kubernetes/` directory and apply the manifests:
```bash
cd kubernetes
kubectl apply -f .
```

This will create the following:
- Deployment with HPA
- HPA for resilience
- Service (NodePort)
- ConfigMap with environment variables
- Secret with the API key

#### Step 4: Expose the Service and Access the Application
To handle potential local network access issues, use Minikube’s built-in service tunnel:
```bash
minikube service stock-ticker-service
```

This should open the service in your default web browser. If the browser doesn’t launch, use the URL provided by the `minikube service` command to access the application. 

#### Step 5: Interact with the Application
Once on the landing page, click the “Get Stock Data” button to retrieve and display the stock information.

### Addressing Limitations
- **NodePort Access**: Minikube’s NodePort might not be accessible directly from the local network. Using `minikube service` as shown in Step 4 should alleviate this issue.
- **Production Considerations**: To improve resilience in production, I would add an Ingress for more accessible routing, add a LoadBalancer for balanced routing, and integrating logging and monitoring tools. Additionally, rate limiting and caching could help mitigate API rate limits.

### Additional Notes
- The `Dockerfile` in the root directory allows you to build the image locally if you’re unable to use the image in the Deployment file. To build the image, navigate to the `Ping` directory and run:
  ```bash
  docker build -t <your-dockerhub-username>/stock_ticker .
  ```

This completes the setup. Please refer to the Kubernetes manifests and Dockerfile or reachout to me at giselle.rodrigues007@gmail.com for assistance with this repo.
