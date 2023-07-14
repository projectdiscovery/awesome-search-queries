<h1 align="center"> Awesome Search Queries </h1>


<h4 align="center"> Community curated list of search queries for various categories across multiple search engines. </h4>

<p align="center">
<a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/license-MIT-_red.svg"></a>
<a href="https://github.com/projectdiscovery/public-bugbounty-programs/issues"><img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat"></a>
<a href="https://twitter.com/pdiscoveryio"><img src="https://img.shields.io/twitter/follow/pdiscoveryio.svg?logo=twitter"></a>
<a href="https://discord.gg/projectdiscovery"><img src="https://img.shields.io/discord/695645237418131507.svg?logo=discord"></a>
</p>

Welcome to Awesome Search Queries, a community-driven project to create an index of search queries for various categories across multiple search engines. This project is inspired by [awesome-hacker-search-engines](https://github.com/edoardottt/awesome-hacker-search-engines) and uses the same search engines supported by [uncover](https://github.com/projectdiscovery/uncover).

This project aims to create a comprehensive, community-curated list of search queries for various categories across multiple search engines. These queries can be used to find instances of specific products, identify potential security vulnerabilities, and gather information for research purposes.

## Understanding the YAML Format

Each product in [QUERIES.yaml](QUERIES.yaml) is represented as a YAML entry. Here's a breakdown of the structure:

```yaml
- name: jira
  type: product
  category: productivity
  engines:
    - platform: shodan
      queries:
        - '"Jira" http.title'
    - platform: censys
      queries:
        - services.software.product=`Jira`
    - platform: fofa
      queries:
        - title="Jira"
    - platform: hunterhow
      queries:
        - product.name=="Jira"
```

- `name`: This is the name of the product. In this case, it's "jira".
- `type`: This is the type of the item. In this case, it's "product". Other possible types could be "vendor", "honeypot", etc.
- `category`: This is optional field to define type category.
- `engines`: This is a list of search engines that can be used to find instances of the product. Each search engine has its own entry in the list.
  - `platform`: This is the name of the search engine.
  - `queries`: This is a list of search queries that can be used on the search engine to find instances of the product. single or multiple queries are supported.


## Contributing

We welcome contributions! There are a couple of ways you can help:

1. **Add New Search Queries for Existing Products:** If you know of a search query that's relevant to an existing product in this project, but isn't currently listed, we'd appreciate if you added it. This will help increase the coverage of this project.

2. **Add New Products with Search Queries:** If there's a product missing in this project that you think should be included, please add it! Be sure to also include a relevant search query for the new product.

See the [CONTRIBUTING.md](CONTRIBUTING.md) file for more information on how to contribute.

Your contributions will be reviewed and, if accepted, merged into the main project. We're looking forward to your input!


## Future Ideas

1. **[uncover](https://github.com/projectdiscovery/uncover) Integration:** support for searching for specific product across all search engines.
2. **[nuclei-templates](https://github.com/projectdiscovery/nuclei-templates) Integration:** auto populate search query metadata in nuclei templates for related products.


Of course! It's important to write clear and understandable documentation. Here's a revised version of your note:

## Important Information

- The file `QUERIES.json` is automatically generated through a GitHub Action workflow. Manual updates to this file are not recommended or necessary.

- The `name` field in all `QUERIES.*` files is formatted to be compatible with the [Common Platform Enumeration](https://csrc.nist.gov/projects/security-content-automation-protocol/specifications/cpe) (CPE) standard. It is highly recommended to maintain this format when available for consistent data representation and to ensure compatibility with systems and services that may use these files.

## Code of Conduct

Please note that this project is released with a Contributor Code of Conduct. By participating in this project you agree to abide by its terms. See the [CODE_OF_CONDUCT.md](https://github.com/projectdiscovery/.github/blob/main/CODE_OF_CONDUCT.md) file for more information.

## License

This project is licensed under the terms of the MIT license. See the [LICENSE](LICENSE.md) file for more information.

-------

<div align="center">

We greatly appreciate your contributions and your efforts in keeping our community dynamic and engaging. ❤️

<a href="https://discord.gg/projectdiscovery"><img src="https://raw.githubusercontent.com/projectdiscovery/nuclei-burp-plugin/main/static/join-discord.png" width="300" alt="Join Discord"></a>

</div>

