import pandas as pd
import seaborn as seas
import matplotlib.pyplot as plt
import matplotlib.patches as mpatches
import random

class PlayerLogout:

    itemFile = "PlayerLogoutInfo.csv"

    def __init__(self):
        self.load_items_from_file(self.itemFile)
        self.all_players_df = pd.DataFrame(columns=['name','current_level','play_time','legacy_count','logout_room','logout_area'])
    
    def load_players_from_file(self,targetFile):
        self.all_players_df = pd.read_csv(targetFile)

    def get_random_color(self, dumpvar):
        return "#"+''.join([random.choice('0123456789ABCDEF') for j in range(6)])
    
manager = PlayerLogout()