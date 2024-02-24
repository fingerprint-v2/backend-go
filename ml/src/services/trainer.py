import random
import time


class Trainer:
    def __init__(self):
        print("Trainer created")

    def train(self, data):
        id_random = random.randint(1000, 99999)
        print(f"Training the model: {id_random}....")
        # print(data)
        for i in range(5):
            print(f"{id_random}: Epoch {i+1}...")
            time.sleep(1)
        print(f"Model trained: {id_random}")
