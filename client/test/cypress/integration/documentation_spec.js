// DONE
describe('Check the documentation page', function () {
  it('All header are visible', function () {
    cy.server()
    cy.route('/api/download/docs').as('docs')
    cy.visit('/documentation')
    cy.wait('@docs')
    cy.get('h1').contains('Installation')
    cy.get('h1').contains('Quick start')
    cy.get('h1').contains('Administration')
    cy.get('h1').contains('Configuration')
    cy.get('h1').contains('API')
    cy.get('h1').contains('Backup & Restore')
    cy.get('h1').contains('Upgrade')
    cy.get('h1').contains('License')
    cy.get('h1').contains('Release Notes')
  })

  it('Payment and email links are set', function () {
    cy.server()
    cy.route('/api/download/docs').as('docs')
    cy.visit('/documentation')
    cy.wait('@docs')
    cy.get('a').contains('PayPal').should('have.attr', 'href', 'https://paypal.me/StanislavValasek')
    cy.get('a').contains('timesheet.simplesw@gmail.com').should('have.attr', 'href', 'mailto:timesheet.simplesw@gmail.com')
  })
})
