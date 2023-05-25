# Hosting and Delivery of Serverless Web Apps on AWS



Two Services

- **S3**
- **Cloudfront**
  - Distrbution: isolated app that you are serving (or any other medium)
  - Origin: where the content is served from
  - Behavior: how cloudfront serves the content (i.e. enable compression)
  - Cache hit and miss: first user will cache miss,  but second user on the same edge location will be cache hit
  - Invalidation: to update the content of distribution 

