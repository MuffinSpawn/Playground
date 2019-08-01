from family_tree.app import create_app


def run_app():
    app = create_app()
    app.run()
app.run(ssl_context='adhoc')

if __name__ == '__main__':
    run_app()

