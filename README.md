
# M31DB 0.5
### The Project is in its alpha realese so expect bugs in it
****Before using this for commercial purposes or modyfing,using it on your own projects see the [LICENSE](LICENSE)****<br>
M31DB is a blazingly fast and easy-to-use database that is designed to power large-scale data applications. It can Directly be extended Using GO Programming Language or Plugins
## Features

- **🔧Easy to use:** M31DB is designed with simplicity and ease-of-use in mind. Its intuitive interface and powerful API make it easy to integrate into your next database and start storing and retrieving data.

- **💻Concurrent:** M31DB supports concurrent operations, allowing multiple users to read and write to the database simultaneously without blocking or slowing down performance.

- **⚡Blazingly fast:** With M31DB, you can expect lightning-fast data processing speeds, even with large datasets. Its optimized algorithms and advanced caching techniques ensure that your data is always available and accessible.
- **📏Minimal:** At Design it is Minimal but combined with Andromeda SQL Parser(Coming Soon) it can perform powerfull yet fast queries
- **Plugins:(Comming Soon)** M31DB Supports Plugin system so you can extend it easily
- **Embedded:** It is designed to be as small as possible so it can be embedded into many softwares [See The Embedding Tutorial]()

## Getting Started

To get started with M31DB, you can follow the steps below:

1. Install M31DB on your local machine or server.
2. Run `m31 init` to initialize M31DB
3. Run command  `m31 start` to start M31DB Server listening on port 6787
4. Now it is all configured to access it you can make a HTTP request to Server running M31DB with the following parameters:
    - **username**: It will be asked on initialization
    - **password**: It will be also asked on initialization
    - **options**: It will be the query you will run seprated by commas
**For Example:**  ```http:\\db.expample.com:6787?username=xyz&password=123&options=select,exampleproj/expamlerow```

For more detailed instructions and examples, check out the documentation at [M31DB Docs](docs/index.html).

## Contributing

M31DB is an open-source project, and we welcome contributions from the community. If you would like to contribute, please check out our [contributing guidelines](https://github.com/M31DB/contributing) and [code of conduct](code-of-conduct/index.html).
## How it Works
M31DB works in a very simple way. Basically, M31DB is a CLI Tool or a  library to interact with JSON but it hides JSON from you by simplifying it to something that is easier to understand that looks like this (x='y',z='y') yeah this looks like CSV with Keys and Values insted of just Values if you want to know more checkout [How It Works?](https://m31db.github.io/how-it-works.html)
## License

M31DB is licensed under the BSD-3 Clause License. See the [LICENSE](LICENSE) file for details.

