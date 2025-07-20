## Security Practice Docker

- [x] Modify the Dockerfile and makes the images size smaller
- [x] Make container run rootles
- [x] Scan images to find vulnerabilities

### How to scan

```shell
$ docker run aquasec/trivy
or
$ winget install AquaSecurity.Trivy
```

```shell
trivy image [YOUR IMAGE]
```
