from locust import HttpLocust, TaskSet, task, between

class Client(TaskSet):

    @task
    def hello(l):
        l.client.get("/hello")

class Test(HttpLocust):
    task_set = Client
    wait_time = between(5.0, 9.0)
