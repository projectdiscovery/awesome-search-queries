## How to Contribute

Contributions to this project are welcome and appreciated. You can contribute by adding new search queries for existing products or by adding new products altogether. Here's how you can do it:

1. **Fork the Repository**: The first step to contributing is to fork the repository to your own GitHub account. This creates your own copy of the entire project, which you can edit as you see fit.

2. **Clone the Repository**: After forking, you'll want to clone the repository to your local machine. This allows you to make changes to the files on your computer.

3. **Find the Relevant YAML File**: Inside this file, you'll find a list of products and associated search queries for each supported search engine. 

4. **Make Your Changes**: Add your new search query to the appropriate place in the YAML file. If you're adding a new product, you'll need to create a new entry in the file. Here's an example structure:

```yaml
- name: product_name
  type: product #product, honeypot, service etc.
  engines:
    - platform: search_engine_name # shodan, censys, fofa, hunter, quake, zoomeye, netlas, criminalip, publicwww, hunterhow, google
      queries:
        - 'your_search_query'
        - 'your_search_query_2'
```

5. **Commit and Push Your Changes**: Once you've made your changes, you'll need to commit them to your local repository and then push them to your forked repository on GitHub.

6. **Create a Pull Request**: Finally, you'll need to create a pull request to have your changes merged into the main project. This is essentially a request to the project maintainer to pull your changes into the main repository.

---

This guide assumes a basic familiarity with Git and GitHub. If you're new to these tools, there are many great resources available online to help you get started.
