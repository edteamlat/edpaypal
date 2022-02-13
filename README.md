# PayPal Integration

Curso de integración con PayPal.

Copie y peque el contenido del archivo `.env.example` a `.env` y coloque los valores de la configuración de su dashboard de PayPal.

```bash
cp .env.example .env
```

Cree los certificados autofirmados para poder ejecutar su programa de manera local:

```bash
mkdir certificates
cd certificates
openssl req -x509 -newkey rsa:4096 -keyout key.pem -nodes -out cert.pem -days 365
```
