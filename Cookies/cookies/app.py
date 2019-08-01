from flask import Flask
# import flask_cors
import hashlib

from cookies.views import health_check, login

# cors = flask_cors.CORS()


# def create_app(config=None, db_path=None):
#     app = Flask(__name__)
#     if config:
#         app.config.update(config)
#     # cors.init_app(app)

#     # app.config['db'] = Database(path=db_path)

#     app.register_blueprint(health_check.blueprint, url_prefix='/api')
#     app.register_blueprint(login.blueprint, url_prefix='/api')
#     # print(app.url_map)

#     return app

def main(config=None):
    app = Flask(__name__)
    if config:
        app.config.update(config)
    app.config.update(
        SESSION_COOKIE_SECURE=True,
        SESSION_COOKIE_HTTPONLY=True,
        SESSION_COOKIE_SAMESITE='Strict',
        PERMANENT_SESSION_LIFETIME=60*60*24*30,
    )
    app.secret_key = hashlib.md5('wiggles'.encode('U8')).hexdigest()
    app.register_blueprint(health_check.blueprint, url_prefix='/api')
    app.register_blueprint(login.blueprint, url_prefix='/api')
    # print(app.url_map)

    app.run(ssl_context='adhoc')


if __name__ == '__main__':
    main()