class Predictor:
    def __init__(self, model=None):
        self.model = model

    def getCurrentModel(self):
        return self.model

    def predict(self, X):
        return "Predicted values"

    def loadModel(self, path):
        self.model = f"Model loaded from {path}"
        return None
