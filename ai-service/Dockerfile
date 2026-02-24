 # Use the latest stable release of python slim as the base image
 FROM python:3.14-slim

# Set Environment Variables
ENV PYTHONDONTWRITEBYTECODE=1 \
    PYTHONUNBUFFERED=1

 # Set working directory inside the container
WORKDIR /app

# Copy the requirements.txt files for dependency installation
COPY requirements.txt ./

# Install Python dependencies
# The --no-cache-dir flag helps reduce the image size
RUN pip install --no-cache-dir -r requirements.txt

# Copy the rest of the application code into the container
COPY . .

# Expose the application's port
EXPOSE 8001

# Run the application
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8001"]