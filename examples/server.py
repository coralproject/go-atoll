from atoll.service import create_app
from atoll import Pipe, Pipeline, register_pipeline


class FooPipe(Pipe):
    input = {
        'body': str,
        'user': str
    }
    output = int

    def __call__(self, input):
        return len(input)

pipeline = Pipeline([FooPipe()], name='score post')

register_pipeline('/score_post', pipeline)

app = create_app()


if __name__ == '__main__':
    app.run(debug=True, port=5001)

