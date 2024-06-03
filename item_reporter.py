import pandas as pd
import seaborn as seas
import matplotlib.pyplot as plt
import matplotlib.patches as mpatches
import random

class ItemReporter:

    itemFile = "AllItems.json"

    def __init__(self):
        self.load_items_from_file(self.itemFile)
    
    def load_items_from_file(self,targetFile):
        self.all_items_df = pd.read_json(targetFile)

    def get_random_color(self, dumpvar):
        return "#"+''.join([random.choice('0123456789ABCDEF') for j in range(6)])
    
    def graph_aggregate_weapon_data(self):
        weapon_df = self.all_items_df.loc[self.all_items_df["type"].str.contains('weapon')].reset_index()
        has_flags_df = weapon_df.loc[weapon_df["flags"].notna(), ["vnum","flags"]]
        #print(weapon_df[["vnum","cost","level_req "]].to_string())
        #weapons per area
        
        #weapons with upgrades
        can_upgrade_mask = self.all_items_df["upgrade_vnum"].ge(1)
        upgradeable_vnums = self.all_items_df.loc[can_upgrade_mask]["vnum"]
        upgrade_vnums = self.all_items_df.loc[can_upgrade_mask].reset_index()["upgrade_vnum"]
        upgradeable_counter = upgradeable_vnums.count()
        print (upgrade_vnums)
        print(f"Number of Upgradeable Weapons: {upgradeable_counter}")

        #weapons that -are- upgrades
        is_upgrade_mask = self.all_items_df["vnum"].apply(lambda vnum: vnum in upgrade_vnums)
        #is_upgrade_mask = self.all_items_df["vnum"].isin(upgrade_vnums)
        are_upgrades_counter = self.all_items_df.loc[is_upgrade_mask]["vnum"].count()
        print(f"Items that are upgrades: {are_upgrades_counter}")

        #final_upgrades_count = self.all_items_df.loc[self.all_items_df]
        #final_upgrade_item_mask = self.all_items_df[["vnum","upgrade_vnum"]].apply(lambda vnum, upgrade_vnum: vnum in upgrade_vnums and upgrade_vnum == 0)
        #final_upgrade_item_list = self.all_items_df[is_upgrade_mask].isin(~upgradeable_vnums)
        
        #final_upgrade_item_list = self.all_items_df[is_upgrade_mask]
        final_upgrade_item_list = self.all_items_df.loc[self.all_items_df["vnum"].isin(upgrade_vnums)]
        final_upgrade_item_list = final_upgrade_item_list.loc[final_upgrade_item_list["upgrade_vnum"].apply(lambda upgrade_vnum: upgrade_vnum == 0)]
        final_upgrade_item_count = final_upgrade_item_list["vnum"].count()
        print(f"Final Upgrade Count: {final_upgrade_item_count}")
        print(final_upgrade_item_list[["vnum","short_desc"]].to_string())
        
        #legendary weapons (possibly with legendaries that aren't upgrades colored diff) 


        #weapons that have affects
        weapons_with_affects = self.all_items_df.loc[self.all_items_df["affects"].notna()]
        weapons_with_affects_count = weapons_with_affects["vnum"].count()
        print(f"Weapons with affects count: {weapons_with_affects_count}")


        #weapons with level limits
        #print(self.all_items_df.loc[self.all_items_df["level_req"].ge(1)], ["vnum","level_req"])
        level_limit_count = self.all_items_df.loc[self.all_items_df["level_req"].ge(1)]["vnum"].count()
        print(f"Level Limits: {level_limit_count}")
        
        #find the number of items that have bind flags
        bind_mask = has_flags_df["flags"].apply(lambda fray: "binding" in fray)
        bind_count = has_flags_df[bind_mask]["vnum"].count()
        print(f"Bind Counter: {bind_count}")
    
    def graph_weapon_damages(self):
        #max_weapon_damage = 0
        weapons = {}
        #print(self.all_items_df["values"].str[1])
        weapon_df = self.all_items_df.loc[self.all_items_df["type"].str.contains('weapon'),["vnum","short_desc","values"]].reset_index()
        #weapon_df[['cond','low_dam', 'high_dam']] = weapon_df['values'].str.split(' ', expand=True)
        weapon_df["low_dam"] = weapon_df["values"].str[1]
        weapon_df["high_dam"] = weapon_df["values"].str[2]*weapon_df["values"].str[1]
        weapon_df["bar_color"] = weapon_df["values"].apply(self.get_random_color)
        
        max_weapon_damage = weapon_df["high_dam"].max()
        #weapon_df.rename(columns={"v1":"low_dam","v2":"high_dam"})
        #weapon_df["high_dam"] = weapon_df.low_dam*weapon_df.high_dam
        print(weapon_df)
        #for item in self.all_items_df:
            #if item['type'] == "weapon":
            
            #print(item)
        
        fig = plt.gcf()
        ax = plt.gcf().subplots()
        #ax = weapon_df.plot.barh()
        #ax.xaxis.grid(True,color="black")
        ax.set_axisbelow(True)
        ax.xaxis.grid(color="black")

        #index = np.arange(len(areas))
        wide = 0.25
        #color_select = {"good":"green","moderate":"yellow","bad":"red","no-data":"gray"}


        #for index in range(len(areas.keys())):
        rand_color = "#"+''.join([random.choice('0123456789ABCDEF') for j in range(6)])
        #    damage_range = area_highs[index] - area_lows[index]
        #    color_choice = color_select[area_quality[index]]
        #    plt.barh(area_names[index], damage_range, color=color_choice, left=area_lows[index], label=area_names[index])
        plt.barh(weapon_df["short_desc"],weapon_df["high_dam"]-weapon_df["low_dam"],color=weapon_df["bar_color"],left=weapon_df["low_dam"],label=weapon_df["short_desc"])
        plt.title("Weapon Damages")
        plt.xlabel("Damage")
        plt.ylabel("Weapon")
        plt.xlim(0,max_weapon_damage)
        plt.xticks(range(0,max_weapon_damage+10,5))



        #legend_handles = [mpatches.Patch(color=value,label=key) for key,value in color_select.items()]
        #legend_handles = [color for color in color_select.values()]
        #legend_labels = [quality for quality in color_select.keys()]
        #plt.legend(handles=legend_handles, loc=4)



        #plt.legend(handles=legend_handles, labels=legend_labels, loc=4)
        #test_lege = mpatches.Patch(color="green",label = "good")
        #plt.legend(handles=[test_lege])
        
        plt.show()


reporter = ItemReporter()
#reporter.graph_weapon_damages()
reporter.graph_aggregate_weapon_data()