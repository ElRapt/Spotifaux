from flask import request, jsonify
import yaml


def content_negotiation(out_body, out_code):
    accepted_type = request.accept_mimetypes.best_match(['application/json', 'application/yaml'])
    if accepted_type == 'application/yaml':
        yaml_response = yaml.dump(out_body, default_flow_style=False)
        return yaml_response, out_code, {'Content-Type': 'application/yaml'}
    else:
        # Default to JSON
        return jsonify(out_body), out_code, {'Content-Type': 'application/json'}
