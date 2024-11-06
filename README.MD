SSO:
    - AUTH (аутентификация)
    - PERMISSIONS (авторизация)
    - USER INFO (UserInfo)


Layers:
    - Transport:
        gRPC Server
    - Service:
        Auth
        Permissions
        EventSender
    - Data:
        PostgreSQL

APP:
    Run();
    Stop();
