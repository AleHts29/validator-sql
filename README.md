# Estructura de proyectos en Golang

```text
- cmd/
    - your_main_package/
        - main.go
- configs/
- docs/
- internal/
    - domain/
        - producto.go
- pkg/
    - validator/
        - validator.go
        - gorm/
            - gorm_validator.go
```

### Detalles de la Estructura
- **cmd/**: Contiene los archivos principales de la aplicación.
    - **tu_paquete_principal/**: Directorio específico de la aplicación que contiene main.go.
- **configs/**: Contiene archivos de configuración si son necesarios.
- **docs/**: Documentación del proyecto si la tienes.
- **internal/**: Contiene los paquetes internos de la aplicación.
    - **domain/**: Contiene las definiciones de las estructuras de dominio, como producto.go.
- **pkg/**: Contiene los paquetes reutilizables que pueden ser importados por otras partes de la aplicación.
    - **validator/**: Paquete que contiene la librería de validación de campos nulos.
        - **validator.go**: Contiene la interfaz y la estructura principal de la librería.
        - **gorm/**: Subpaquete que proporciona implementaciones específicas de Gorm para la validación de campos nulos.

## Uso
Para utilizar la librería de validación de campos nulos en tu aplicación, simplemente importa el paquete `validator` en cualquier parte donde necesites realizar validaciones de campos nulos.

```go
import (
    "tu_paquete/pkg/validator"
)
