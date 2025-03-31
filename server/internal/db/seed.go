package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"

	"github.com/iykeevans/go-social/server/internal/store"
)

var usernames = []string{
	"alice", "bob", "charlie", "dave", "eve", "frank", "grace", "heidi", "ivan", "judy",
	"ken", "linda", "mallory", "nancy", "oscar", "peggy", "quentin", "rachel", "steve", "trudy",
	"ursula", "victor", "wendy", "xander", "yvonne", "zack", "amber", "brian", "chloe", "daniel",
	"elena", "felix", "gina", "harry", "isla", "jack", "karen", "leo", "mia", "nathan",
	"olivia", "paul", "quincy", "rebecca", "sam", "tina", "ulysses", "vanessa", "will", "zoe",
}

var titles = []string{
	"10 Tips for Learning Go Faster",
	"The Future of Web Development in 2025",
	"How to Build a REST API with Node.js",
	"Why You Should Care About Clean Code",
	"React vs Vue: Which is Better for Your Next Project?",
	"A Beginner’s Guide to TypeScript",
	"What Makes a Great Backend Developer?",
	"How to Get Started with DevOps",
	"5 Tools Every Full-Stack Developer Needs",
	"Mastering Asynchronous JavaScript",
	"The Power of Microservices in Modern Apps",
	"What I Learned After Switching to Go",
	"How to Handle Errors in Node.js",
	"Making Your Web App Scalable",
	"Exploring the World of Serverless Architectures",
	"Frontend Frameworks: Choosing the Right One",
	"Building Secure Web Applications from the Ground Up",
	"The Importance of Unit Testing",
	"CI/CD in the Real World: A Practical Approach",
	"How to Optimize Your Web App for Speed",
}

var contents = []string{
	"Exploring the latest trends in AI development and their real-world applications.",
	"A step-by-step guide to deploying your first application on AWS.",
	"Why coding standards matter and how to implement them in your team.",
	"The ultimate comparison: MongoDB vs PostgreSQL for backend development.",
	"How I transitioned from frontend to full-stack development and what I learned.",
	"The pros and cons of using serverless architectures in modern apps.",
	"A beginner's guide to containerization and Docker for developers.",
	"Effective debugging techniques every developer should know.",
	"Building a real-time chat application with WebSockets and Node.js.",
	"Everything you need to know about JavaScript's 'this' keyword.",
	"Exploring the benefits of test-driven development (TDD) and how to get started.",
	"Why performance optimization is crucial for user experience and how to achieve it.",
	"Introduction to GraphQL: The future of querying APIs.",
	"The importance of version control and how to use Git like a pro.",
	"How to build a simple, yet effective, to-do list app with React.",
	"What are design patterns, and how can they help solve complex problems?",
	"A guide to continuous integration and delivery (CI/CD) with GitHub Actions.",
	"Building scalable web apps with microservices architecture.",
	"Understanding OAuth and its role in modern authentication.",
	"Exploring WebAssembly and its potential to revolutionize the web.",
}

var tags = []string{
	"GoLang", "web development", "nodejs", "backend", "frontend",
	"React", "TypeScript", "microservices", "serverless", "AWS",
	"Docker", "CI/CD", "unit testing", "performance optimization", "DevOps",
	"GraphQL", "MongoDB", "PostgreSQL", "authenticating users", "cloud computing",
}

var fakeComments = []string{
	"Great article! I learned a lot about Go.",
	"This is exactly what I needed. Thanks for the detailed guide.",
	"I love the clarity in this post. Very helpful!",
	"The code examples are easy to follow. Keep it up!",
	"Nice post, but I’d love to see more examples for beginners.",
	"Very insightful! You’ve cleared up some doubts I had.",
	"Great post! Any chance you could cover X in a future post?",
	"I appreciate the deep dive into serverless architecture.",
	"This helped me understand TypeScript much better. Thanks!",
	"Could you explain more about handling errors in Node.js?",
	"This is a fantastic resource! I’ll definitely share it with my team.",
	"I’ve been looking for a good guide on CI/CD. This was perfect!",
	"The section on performance optimization was super useful.",
	"I never knew about the power of GraphQL. This was eye-opening.",
	"Could you share more about the MongoDB vs PostgreSQL comparison?",
	"Nice job explaining complex concepts in simple terms!",
	"I tried your approach to building a chat app, and it worked great!",
	"This post motivated me to start using Docker in my workflow.",
	"This is exactly what I was struggling with, thanks for the solution!",
	"Awesome! Looking forward to more posts on cloud computing.",
}

func Seed(store store.Storage, db *sql.DB) {
	ctx := context.Background()

	users := generateUsers(100)
	tx, _ := db.BeginTx(ctx, nil)

	for _, user := range users {
		if err := store.Users.Create(ctx, tx, user); err != nil {
			_ = tx.Rollback()
			log.Println("error creating user:", err)
			return
		}
	}

	posts := generatePosts(200, users)
	for _, post := range posts {
		if err := store.Posts.Create(ctx, post); err != nil {
			log.Println("error creating post:", err)
			return
		}
	}

	comments := generateComments(200, users, posts)
	for _, comment := range comments {
		if err := store.Comments.Create(ctx, comment); err != nil {
			log.Println("error creating post:", err)
			return
		}
	}

	log.Println("seeding complete")
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := 0; i < num; i++ {
		users[i] = &store.User{
			Username: usernames[i%len(usernames)] + fmt.Sprintf("%d", i),
			Email:    usernames[i%len(usernames)] + fmt.Sprintf("%d", i) + "@example.com",
		}
	}

	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]

		posts[i] = &store.Post{
			UserID:  user.ID,
			Title:   titles[rand.Intn(len(titles))],
			Content: contents[rand.Intn(len(contents))],
			Tags: []string{
				tags[rand.Intn(len(tags))],
				tags[rand.Intn(len(tags))],
			},
		}
	}

	return posts
}

func generateComments(num int, users []*store.User, posts []*store.Post) []*store.Comment {
	comments := make([]*store.Comment, num)

	for i := 0; i < num; i++ {
		user := users[rand.Intn(len(users))]
		post := posts[rand.Intn(len(posts))]

		comments[i] = &store.Comment{
			UserID:  user.ID,
			PostID:  post.ID,
			Content: fakeComments[rand.Intn(len(fakeComments))],
		}
	}

	return comments
}
