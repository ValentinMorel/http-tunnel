# Http tunnel to ssh

# Usage 

Listen ssh connections:

```sh
go run main.go
```

Connect with ssh: 

```sh
ssh localhost -p 2222
```

output : 

```
tunnel ID ->  1200233730143388165
```

Access it : 

```
http://localhost:3000/list?id=1200233730143388165&dir=../../../
```

Should give: 

![](./docs/result.png)


# Workflow design

```mermaid

    flowchart TD
    A[Start] --> B[HTTP Server]
    B --> |"2"| C[Handle Requests]
    C --> D{Request Type}
    D --> |"List Files"| E[List Directory]
    D --> |"Download File"| F[Serve File]
    
    E --> G[Display File List]
    F --> H{File Type}
    H --> |"Display"| I[Inline Display]
    H --> |"Download"| J[Attachment Download]
    
    B --> |"1"| K[SSH Server]
    K --> L[Handle SSH Connections]
    L --> M[Establish Tunnel]
    M --> N[Manage Tunnel]
    
    %% File Handling
    F --> O{File Extension}
    O --> |".go"| P[Detect MIME Type]
    O --> |".txt"| P
    O --> |".html"| P
    O --> |".pdf"| P
    O --> |"Unknown"| Q[Read File Content]
    P --> R[Serve File Inline or as Attachment]
    Q --> R
    
    %% Client Interaction
    I --> S[Display in Browser]
    J --> T[Prompt Download]
    
    %% End
    S --> U[End]
    T --> U
    U --> A

```