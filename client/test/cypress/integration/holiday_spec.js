// DONE
describe('Holiday page - has header and > 5 rows', function () {
  it('Table has a correct header', function () {
    cy.server()
    cy.route('/api/holidays').as('holidays')
    cy.visit('/holidays')
    cy.wait('@holidays')
    cy.get('th').contains('Date')
    cy.get('th').contains('Day')
    cy.get('th').contains('Description')
  })

  it('Table contains at least 5 rows', function () {
    cy.server()
    cy.route('/api/holidays').as('holidays')
    cy.visit('/holidays')
    cy.wait('@holidays')
    cy.get('tbody').children().should('have.length.at.least', 5)
  })
})
