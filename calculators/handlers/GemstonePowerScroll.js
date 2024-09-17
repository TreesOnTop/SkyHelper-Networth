const { APPLICATION_WORTH } = require('../../constants/applicationWorth');

/**
 * A handler for the Gemstone Power Scroll modifier on an item.
 */
class GemstonePowerScrollHandler {
    /**
     * Checks if the handler applies to the item
     * @param {object} item The item data
     * @returns {boolean} Whether the handler applies to the item
     */
    applies(item) {
        return item.extraAttributes.power_ability_scroll;
    }

    /**
     * Calculates and adds the price of the modifier to the item
     * @param {object} item The item data
     * @param {object} prices A prices object generated from the getPrices function
     */
    calculate(item, prices) {
        // TODO: Remove This
        // ? NOTE: USE THIS IF YOU'RE COMPARING OLD AND NEW
        // ?
        // ? (COMMENT OUT THIS PART, IT WAS FULLY BROKEN IN THE OLD CALCULATION, IT WOULD ALWAYS RETURN 0)
        // ?

        const calculationData = {
            id: item.extraAttributes.power_ability_scroll,
            type: 'GEMSTONE_POWER_SCROLL',
            price: (prices[item.extraAttributes.power_ability_scroll] || 0) * APPLICATION_WORTH.gemstonePowerScroll,
            count: 1,
        };
        item.price += calculationData.price;
        item.calculation.push(calculationData);
    }
}

module.exports = GemstonePowerScrollHandler;