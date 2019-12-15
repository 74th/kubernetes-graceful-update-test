from invoke import task

@task
def run_localhost(c):
    c.run("locust -H http://localhost:8080")
