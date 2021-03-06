# chocolate-waffle

## Overview
chocolate-waffle aims to clone shorby.com for education purpose. chocolate-waffle is still in very early state and under development. If you want to contribute please [Open Issue](https://github.com/billysutomo/chocolate-waffle/issues/new/choose) or Merge Request.

![UI Image](ui.png)

## Quick Start
### Service
* This project use migration tools dbmate. Please install dbmate https://github.com/amacneil/dbmate
* Run `dbmate migrate` to create database
* run `make devservice` to run service API
### Web
* run `make devweb` to run web

## Tech Stack
### UI
* TypeScript
* React
* Redux
* styled-component
### Service
* Golang
* Postgre

## ERD

![ERD Image](erd.png)
### List Table
#### users
Manage users data and credential
* id: primary key
* name: user name
* password: user password
#### projects
each user can have one or more projects
* id: primary key
* id_user: user foreign key, define which user project belonging
* url: unique url for this project
* profile_picture: profile picture url path 
* title: project title
* description: project description
#### elements
Each project can have one or more elements. 
* id: primary key
* id_project: project foreign key, define which project element belonging
* ordernum: element ordering number from top to bottom
* type: element type, [messenger, block, social_link]
* body: json information, define by view

