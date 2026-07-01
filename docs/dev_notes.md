# Lessons Learned

1. Every handler should validate the HTTP method before processing the request.

2. JSON decoding can fail, so every request body must be checked.

3. HTTP handlers should call the storage layer instead of manipulating data directly.

4. Tests should verify behavior, not just status codes.

5. Small commits make debugging much easier than one large commit.

6. Query parameters are the standard way to pass identifiers in GET and DELETE requests.

7. Keeping handlers small makes the code easier to understand and test.