# Usage


```yaml
pipeline:
  rollout:
    image: mikeifomin/drone-swarmupd
    settings:
      url: 
        from_secret: swarmupd_url
      token:
        from_secret: swarmupd_token
      service_name: api_prod
      new_tag: ${DRONE_COMMIT_SHA:0:10}
```

