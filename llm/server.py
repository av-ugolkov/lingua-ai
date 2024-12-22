import os
from flask import Flask, request, send_file
import torch
from transformers import pipeline
from datasets import load_dataset
import soundfile as sf

app = Flask(__name__)


synthesiser = pipeline("text-to-speech", "./speecht5_tts")
embeddings_dataset = load_dataset("Matthijs/cmu-arctic-xvectors", split="validation")
speaker_embedding = torch.tensor(embeddings_dataset[7306]["xvector"]).unsqueeze(0)
# You can replace this embedding with your own as well.

@app.route('/synthesize', methods=['POST'])
def synthesize():
    data = request.json
    text = data.get('text','')
    name = data.get('name','')

    if not text:
        return {"error": "Text is required"}, 400
    if not name:
        return {"error": "Name is required"}, 400

    speech = synthesiser(text, forward_params={"speaker_embeddings": speaker_embedding})
    sf.write("{name}.wav", speech["audio"], samplerate=speech["sampling_rate"])
    
    return send_file("{name}.wav", mimetype='audio/wav')

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=7000)
