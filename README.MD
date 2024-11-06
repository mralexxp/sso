SSO:
    - AUTH (аутентификация)
    - TODO: PERMISSIONS (авторизация)
    - TODO: USER INFO (UserInfo)


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
