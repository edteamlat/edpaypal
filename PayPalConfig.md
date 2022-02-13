# Configuración PayPal

En esta sección veremos los pasos iniciales para configurar PayPal.

## 1. Crear una cuenta en PayPal

Debes ingresar a [paypal.com](http://paypal.com) y registrarte. Puedes registrarte como `Personal` para realizar toda la integración que vamos a ver en el curso. Pero para vender si es mejor que te registres como `Negocios`.

## 2. Ingresar al dashboard de desarrollador

Una vez hayas creado tu cuenta en [paypal.com](http://paypal.com) debes ingresar al dashboard de programador en el que podrás realizar todos los pasos necesarios para la integración de PayPal. Para esto debes ingresar a [developer.paypal.com](http://developer.paypal.com) y dar clic en el botón `Log into Dashboard`.

## 3. Accounts

En esta pestaña encontraremos las dos cuentas que PayPal nos entrega por defecto, una de negocio y una personal. Muchos desarrolladores prefieren crear cuentas adicionales a estas que son creadas por PayPal. Es tu decisión si creas nuevas o no. Nosotros no vamos a crear nuevas cuentas. Lo que sí haremos es copiarlas y pegarlas en un sitio que sea de fácil acceso ya que buscar la información de las cuentas desde esta pestaña puede ser tedioso.

En cada cuenta iremos a la columna `Manage accounts` y daremos clic en los 3 puntos, luego seleccionamos `view/edit account` y nos permite ver la contraseña. También puedes cambiarla si quieres. Guardaremos los datos en un sitio de fácil acceso.

- Datos

  EMPRESA: [sb-hwcnv11885178@business.example.com](mailto:sb-hwcnv11885178@business.example.com)

  CLAVE: 4(R2]j^m

  PERSONAL: [sb-1c32211889669@personal.example.com](mailto:sb-1c32211889669@personal.example.com)

  CLAVE: L>%Ztad6

  ClientID: AdomWD1SSXuqFpXcUfgMxYcQShoCUoINeBCnis7jNEfJJpbFfIW555tY9b6eX7pDlyiQB40WEk5AWGyp


## 4. My Apps & Credentials

La primera sección que encontraremos en el dashboard de desarrollador es la de `My apps & credentials`, aquí podremos crear las apps que queramos para hacer pruebas. Utilizaremos el entorno `sandbox` el cual nos permite hacer las pruebas necesarias para la integración. Puedes utilizar la aplicación que viene por default, o puedes crear una nueva. Nosotros crearemos una nueva.

Daremos clic en el botón `Create app` que nos llevará a la nueva ventana, en esa ventana daremos un nombre a nuestra app, seleccionaremos el tipo de app `Merchant` (la otra opción es para servir como plataforma de vendedores) y seleccionaremos la cuenta de negocio que viene por defecto o la que hayas creado.

Por último podemos agregar un webhook listener de nuestro servidor. Esto lo podemos hacer más adelante.

## 5. Notifications

En esta pestaña podrás ver cómo se envían los correos tanto al negocio como al comprador.

## 6. API calls

En esta pestaña podrás ver las llamadas a las API de PayPal y su estado (correcto o error).

## 7. Webhooks Events

En esta pestaña veremos los eventos enviados a nuestro servidor y podremos reenviar los eventos si es que por algún motivo no los recibimos o los procesamos incorrectamente.

## 8. Mock / Webhook simulator

Esta pestaña nos permite enviar eventos simulados hacia nuestro servidor. Es muy útil para ver el contenido de todos los eventos. Incluso no necesitas un servidor funcionando ya que cuando se envía nos mostrará la data enviada y podremos hacer un mock con esa información.

Listo, ya tenemos lista la configuración para comenzar a desarrollar la integración con PayPal.