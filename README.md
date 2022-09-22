# Monitoring App Written in GoLang

This project is the final project of Snapp Summer course, with mentoring of [Parham Alavani](https://github.com/1995parham). 

And is developed by [Ghazal](https://github.com/ghazallghe) and [Mohammad](https://github.com/MohammadShabaniSBU).

## Introduction

As a developer we want to monitor some web application for several reason. For example is our web application working correct? Or using a third party application is a good decision or not? For this purpose we need an application to make requests to our targeted urls, in a regular time interval and store the results. So then the user can observe how the application works.

## Implementation

First of all we need users to be authenticated. Because it is RestApi application we choose to generate JWT tokens to identify users. Secondly we need each user to define her/his own urls which she/he wants to be monitored. At this point we run our scheduling system which makes requests to all the urls we stored, in a specific time interval and store their response results. There are `statistics` and `alerts` endpoints which are used to observe these results.



## Trade offs

#### Fiber vs Other libraries in go

We chose Fiber because of its performance and its structure which is inspired from Express framework and It has a better documentation than other libraries.

#### SQL vs NoSQL

Because our resources for this application (users, urls and statistics) have schemas and those will not change during development we prefer to choose SQL over NoSQL. Also SQL is faster.

#### MySql vs PostgreSql

When we test the urls in parallel and get the result we need to write them in database and because we do this in parallel we need a database to handle the concurrency better. So we chose PostgreSQL.

#### Postman vs Swagger for API documentation

Postman environment is very easier to install, set up and use compared with swagger and it also meets our needs for integration tests and etc.
[documentation](https://documenter.getpostman.com/view/18401078/)

## How to install

**Requirements**

+ docker
+ docker-compose

**How to run**

1. Clone this repo
2. Copy `.env.example` to `.env`. (It has some defaults value to work out of the box but you can change any of them)
3. Run
   + For development `docker-compose up`
   + For production `docker-compose -f docker-compose.production.yml up -d`


