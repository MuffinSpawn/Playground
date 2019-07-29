from flask import blueprints, jsonify, make_response, request, session
import hashlib
import http.cookies

blueprint = blueprints.Blueprint('login', __name__)


@blueprint.route('login/', methods=['GET'])
def login():
    # http.cookies.Morsel._reserved['samesite'] = 'SameSite'
    # cookies = http.cookies.SimpleCookie()
    # cookies['session'] = hashlib.md5('wiggles'.encode('U8')).hexdigest()
    # cookies['session']['Path'] = '/'
    # cookies['session']['Max-Age'] = 60
    # cookies['session']['Secure'] = True
    # cookies['session']['HttpOnly'] = True
    # cookies['session']['SameSite'] = 'Strict'
    # resp = make_response(render_template(...))

    # print(request.cookies)
    response_value = {}
    if 'session' in request.cookies:
        response_value['User ID'] = session['user_id']
        response_value['permenant'] = session.permanent
    resp = make_response(jsonify(response_value))
    session.clear()
    session['user_id'] = '42'
    session.permanent = True
    
    return resp 

