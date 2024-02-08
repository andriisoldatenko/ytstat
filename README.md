# YouTube stats tool

### Intro

It's my personal tool to get all videos from given playlist,
usually playlist includes all videos from an conference 
and select some of them and watch later.
So I can filter them by `viewCount` (default ordering).

#### Install:

```bash
go install github.com/andriisoldatenko/ytstat@latest
```

#### Usage:
```
ytstat -playListID PLHhKcdBlprMdIMzUZX6ho0OPTikTamLwa -apiKey <apiKey> | jq
```

sample output:

```json
...
  {
    "title": "Hands-On with an Envoy API Gateway, Now with GraphQL! - Jim Barton, Solo.io",
    "viewCount": 352
  },
  {
    "title": "Container-native pipelines with Drone CI - Jim Sheldon, Harness",
    "viewCount": 427
  },
  {
    "title": "Hobbyfarm - an OpenSource Kubernetes Training Environment - Enrico Bartz, SVA",
    "viewCount": 506
  },
  {
    "title": "IPv6 / Dual-Stack in Kubernetes - Why, When, Where and How? - Rastislav Szabo, Kubermatic",
    "viewCount": 697
  },
  {
    "title": "ContainerDays 2022 Aftermovie",
    "viewCount": 857
  },
  {
    "title": "The future of CRDs in a post-cluster world - Sebastian Scheele & Stefan Schimanski",
    "viewCount": 957
  }
]
```


### TODO:
- add more stats